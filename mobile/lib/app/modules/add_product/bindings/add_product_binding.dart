import 'package:demo_store/app/modules/add_product/controllers/storage_controller.dart';
import 'package:get/get.dart';

import '../controllers/add_product_controller.dart';

class AddProductBinding extends Bindings {
  @override
  void dependencies() {
    Get.lazyPut<AddProductController>(
      () => AddProductController(Get.find()),
    );
    Get.lazyPut<StorageController>(
      () => StorageController(Get.find()),
    );
  }
}
