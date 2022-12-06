import 'package:demo_store/app/routes/app_pages.dart';
import 'package:demo_store/internal/di/di.dart';
import 'package:demo_store/internal/repository/notification.dart';
import 'package:demo_store/internal/repository/product.dart';
import 'package:demo_store/internal/repository/user.dart';
import 'package:flutter/widgets.dart';
import 'package:flutter_localizations/flutter_localizations.dart';
import 'package:get/instance_manager.dart';
import 'package:get/route_manager.dart';
import 'package:kiwi/kiwi.dart';
import 'package:sizer/sizer.dart';

import 'resource/colors.dart';

void main() async {
  WidgetsFlutterBinding.ensureInitialized();
  await initDI();
  runApp(
    Sizer(builder: (context, orientation, deviceType) {
      return GetMaterialApp(
        getPages: AppPages.routes,
        initialRoute: AppPages.INITIAL,
        title: 'Demo Store',
        debugShowCheckedModeBanner: false,
        enableLog: true,
        logWriterCallback: Logger.write,
        localizationsDelegates: const [
          GlobalMaterialLocalizations.delegate,
          GlobalWidgetsLocalizations.delegate,
        ],
        theme: createTheme(),
        initialBinding: BindingsBuilder(() {
          Get.put(KiwiContainer().resolve<UserRepository>());
          Get.put(KiwiContainer().resolve<ProductRepository>());
          Get.put(KiwiContainer().resolve<NotificationRepository>());
        }),
        supportedLocales: const [
          Locale('en', 'EN'),
        ],
      );
    }),
  );
}

mixin Logger {
  // Sample of abstract logging function
  static void write(String text, {bool isError = false}) {
    Future.microtask(() => print('** $text. isError: [$isError]'));
  }
}
