import 'package:demo_store/internal/datasources/api.dart';
import 'package:demo_store/internal/repository/notification.dart';
import 'package:demo_store/internal/repository/product.dart';
import 'package:demo_store/internal/repository/user.dart';
import 'package:demo_store/tools/device_name.dart';
import 'package:hive/hive.dart';
import 'package:http/http.dart';
import 'package:kiwi/kiwi.dart';
import 'package:path_provider/path_provider.dart';

Future<void> initDI() async {
  final directory = await getApplicationSupportDirectory();
  Hive.init(directory.path);
  final deviceModel = await deviceName();
  final DataSource datasource =
      DataSource(client: Client(), deviceModel: deviceModel);
  final KiwiContainer container = KiwiContainer();

  final UserRepository userRepo = UserRepository(datasource);
  final ProductRepository productRepo = ProductRepository(datasource);
  final NotificationRepository notificationRepo =
      NotificationRepository(datasource);

  container.registerSingleton((c) => userRepo);
  container.registerSingleton((c) => productRepo);
  container.registerSingleton((c) => notificationRepo);
}
