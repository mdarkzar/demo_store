interface Failure {
  get message(): string;
}

class ResponseFailure implements Failure {
  msg: string = "";

  constructor(msg: string) {
    this.msg = capitalizeFirstLetter(msg);
  }

  get message(): string {
    return this.msg;
  }
}

class ServerFailure implements Failure {
  get message(): string {
    return "Сервер временно недоступен";
  }
}

class AuthFailure implements Failure {
  get message(): string {
    return "Вы не авторизованы";
  }
}

class NoDataFailure implements Failure {
  get message(): string {
    return "Данные не найдены";
  }
}

class MethodNotFoundFailure implements Failure {
  get message(): string {
    return "Вызов неизвестного метода";
  }
}

function capitalizeFirstLetter(s: string) {
  return s.charAt(0).toUpperCase() + s.slice(1);
}

export type { Failure };

export {
  ResponseFailure,
  ServerFailure,
  AuthFailure,
  NoDataFailure,
  MethodNotFoundFailure,
};
