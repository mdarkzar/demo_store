import 'package:demo_store/internal/entity/storage/storage.dart';
import 'package:demo_store/internal/repository/product.dart';
import 'package:demo_store/widgets/notification.dart';
import 'package:get/get.dart';

class StorageController extends GetxController {
  final ProductRepository productRepo;
  StorageController(this.productRepo);

  bool loading = false;
  RxList<Storage>? storageList;

  @override
  void onInit() {
    loadStorageList();
    super.onInit();
  }

  @override
  void onReady() {
    super.onReady();
  }

  loadStorageList() async {
    final r = await productRepo.loadStorageList();
    r.fold((l) => errorMessage(l.message), (r) {
      storageList = RxList(r.storageList!);
    });
    update();
  }

  @override
  void onClose() {}
}
