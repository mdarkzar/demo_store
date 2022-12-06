import 'dart:io';

import 'package:http/http.dart' as http;
import 'package:shared_preferences/shared_preferences.dart';

class HttpWithSession {
  final http.Client client;
  HttpWithSession(this.client);

  Map<String, String> headers = {
    'Content-Type': 'application/json; charset=UTF-8',
  };

  Future<http.Response> get(Uri url) async {
    await _loadCookies();

    http.Response response = await client.get(
      url,
      headers: headers,
    );
    _updateCookie(response);
    return response;
  }

  Future<http.Response> post(Uri url, dynamic data) async {
    await _loadCookies();

    http.Response response =
        await client.post(url, body: data, headers: headers);
    _updateCookie(response);
    return response;
  }

  Future<http.Response> postMultipart(Uri url, File file) async {
    await _loadCookies();

    var request = http.MultipartRequest('POST', url);
    headers.forEach((k, v) {
      request.headers[k] = v;
    });

    request.files.add(http.MultipartFile.fromBytes(
        'file', file.readAsBytesSync(),
        filename: file.path));

    var streamedResponse = await request.send();

    http.Response response = await http.Response.fromStream(streamedResponse);
    return response;
  }

  Future<void> _updateCookie(http.Response response) async {
    String? rawCookie = response.headers['set-cookie'];
    if (rawCookie != null) {
      int index = rawCookie.indexOf(';');
      headers['Cookie'] =
          (index == -1) ? rawCookie : rawCookie.substring(0, index);

      SharedPreferences sharedPreferences =
          await SharedPreferences.getInstance();
      await sharedPreferences.setString("cookie", headers['Cookie'] ?? '');
      return;
    }
  }

  updateHeaders(String key, value) {
    headers[key] = value;
  }

  Future<void> _loadCookies() async {
    SharedPreferences sharedPreferences = await SharedPreferences.getInstance();
    final String? cookie = sharedPreferences.getString("cookie");
    if (cookie != null) {
      headers['Cookie'] = cookie;
    }
  }

  Future<String> cookie() async {
    SharedPreferences sharedPreferences = await SharedPreferences.getInstance();
    final cookie = sharedPreferences.getString("cookie");
    return cookie!;
  }
}
