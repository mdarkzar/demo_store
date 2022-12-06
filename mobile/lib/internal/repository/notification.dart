import 'package:dartz/dartz.dart';
import 'package:demo_store/internal/core/error/failures.dart';
import 'package:demo_store/internal/datasources/api.dart';
import 'package:demo_store/internal/entity/notification/notification.dart';

class NotificationRepository {
  final DataSource source;

  NotificationRepository(this.source);

  Future<Either<Failure, NotificationList>> loadNew() async {
    final res = await source.get("/notification/new");

    return res.fold((l) => Left(l), (r) => Right(NotificationList.fromJson(r)));
  }
}
