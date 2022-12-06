import 'package:equatable/equatable.dart';

abstract class Failure extends Equatable {
  String get message;

  @override
  List<Object> get props => [];
}

// General failures
class ApiFailure extends Failure {
  @override
  get message => "Произошла непредвиденная ошибка";
}

// General failures
class ServerFailure extends Failure {
  @override
  get message => "Нет доступа к серверу";
}

class AuthFailure extends Failure {
  @override
  get message => "Вы не авторизованы";
}

class ResponseFailure extends Failure {
  final String msg;
  ResponseFailure(this.msg);

  @override
  get message => msg;
}

class StorageFailure extends Failure {
  @override
  get message => "Произошла ошибка при записи в локальное хранилище";
}

class NoDataFailure extends Failure {
  @override
  get message => "данные не найдены";
}
