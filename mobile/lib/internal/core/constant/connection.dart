class Connection {
  static const baseURL = "demo-store.darkzar.uz";
  static const devBaseURL = "192.168.5.233";
  static const errNoConnectMessage =
      "Не удается подключиться к серверу. Проверьте соединение с Интернетом или отключите VPN";

  static Uri getUri(String path) {
    // return Uri.https(Connection.baseURL, path);
    return Uri.http(Connection.devBaseURL, path);
  }
}
