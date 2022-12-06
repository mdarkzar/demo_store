import 'package:dartz/dartz.dart';
import 'package:demo_store/internal/core/error/failures.dart';
import 'package:demo_store/internal/datasources/api.dart';
import 'package:demo_store/internal/entity/user/user.dart';
import 'package:get/route_manager.dart';

class UserRepository {
  final DataSource source;

  UserRepository(this.source);

  Future<Either<Failure, void>> login(String login, String password) async {
    final res =
        await source.post("/user/auth", {"login": login, "password": password});

    return res.fold((l) => Left(l), (r) => const Right(null));
  }

  Future<Either<Failure, User>> profile() async {
    final res = await source.get(
      "/user/profile",
    );

    return res.fold((l) => Left(l), (r) => Right(User.fromJson(r['user'])));
  }

  Future<Either<Failure, void>> logout() async {
    final res = await source.post("/user/logout", {});

    return res.fold((l) => Left(l), (r) => const Right(null));
  }
}
