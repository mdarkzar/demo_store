import 'dart:async';
import 'dart:convert';
import 'dart:io';

import 'package:dartz/dartz.dart';
import 'package:demo_store/internal/core/constant/connection.dart';
import 'package:demo_store/internal/core/error/failures.dart';
import 'package:demo_store/internal/datasources/response.dart';
import 'package:demo_store/internal/entity/response/response.dart';
import 'package:demo_store/tools/http_with_session.dart';
import 'package:http/http.dart' as http;
import 'package:retry/retry.dart';

class DataSource {
  final http.Client client;
  final String deviceModel;
  late HttpWithSession session;
  DataSource({required this.client, required this.deviceModel}) {
    session = HttpWithSession(client);
  }

  Future<Either<Failure, dynamic>> post(String path, dynamic params,
      {bool retry = true}) async {
    final response = await _request(
        Connection.getUri("/api/v1$path"), HttpMethod.post, params,
        retry: retry);

    return response.fold((l) => Left(l), (r) => _parseResponse(r));
  }

  Future<Either<Failure, dynamic>> get(String path, {bool retry = true}) async {
    final response = await _request(
        Connection.getUri("/api/v1$path"), HttpMethod.get, null,
        retry: retry);

    return response.fold((l) => Left(l), (r) => _parseResponse(r));
  }

  Future<Either<Failure, ApiResponse>> _request(
      Uri url, HttpMethod method, dynamic params,
      {bool retry = false}) async {
    ApiResponse response;
    session.updateHeaders('User-Agent', deviceModel);
    try {
      const retryOption = RetryOptions(maxAttempts: 8);
      http.Response httpResponse;

      if (retry) {
        httpResponse = await retryOption.retry(
          () => _httpRequest(url, method, params),
          retryIf: (e) => e is SocketException || e is TimeoutException,
        );
      } else {
        httpResponse = await _httpRequest(url, method, params);
      }

      if (httpResponse.statusCode == HttpStatus.ok &&
          httpResponse.body != "null") {
        response = ApiResponse(jsonDecode((httpResponse.body)));
      } else if (httpResponse.body == "null") {
        return Right(ApiResponse({}));
      } else if (httpResponse.statusCode == HttpStatus.unauthorized) {
        return Left(AuthFailure());
      } else {
        return Left(ApiFailure());
      }
    } catch (e) {
      if (e is SocketException || e is TimeoutException) {
        return Left(ServerFailure());
      } else {
        return Left(ResponseFailure(e.toString()));
      }
    }

    return Right(response);
  }

  Either<Failure, dynamic> _parseResponse(ApiResponse resp) {
    final response = Response.fromJson(resp.result);

    if (response.error != null) {
      return Left(ResponseFailure(response.error));
    }

    if (response.result == 'OK') {
      return const Right({});
    }

    return Right(response.result);
  }

  Future<http.Response> _httpRequest(
      Uri url, HttpMethod method, dynamic data) async {
    switch (method) {
      case HttpMethod.post:
        return await session.post(url, jsonEncode(data));
      case HttpMethod.get:
        return await session.get(url);
    }
  }
}
