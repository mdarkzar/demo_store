import 'package:demo_store/app/modules/home/controllers/d_controller.dart';
import 'package:get/get.dart';

import '../controllers/home_controller.dart';

class HomeBinding extends Bindings {
  @override
  void dependencies() {
    Get.lazyPut<HomeController>(
      () => HomeController(Get.find()),
    );
    Get.lazyPut<DController>(
      () => DController(Get.find()),
    );
  }
}
