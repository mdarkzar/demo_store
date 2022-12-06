import 'package:demo_store/app/routes/app_pages.dart';
import 'package:demo_store/internal/repository/product.dart';
import 'package:demo_store/internal/repository/user.dart';
import 'package:demo_store/widgets/notification.dart';
import 'package:get/get.dart';

class HomeController extends GetxController with StateMixin {
  final ProductRepository productRepo;
  HomeController(this.productRepo);

  @override
  void onInit() {
    super.onInit();
    loadData();
  }

  loadData() async {
    change(null, status: RxStatus.loading());
    final r = await productRepo.loadAll();
    r.fold(
        (l) => errorMessage(l.message),
        (r) => change(r,
            status: r.productList == null || r.productList!.isEmpty
                ? RxStatus.empty()
                : RxStatus.success()));
  }

  deleteProduct(int productID) async {
    change(null, status: RxStatus.loading());
    final r = await productRepo.delete(productID);
    r.fold((l) => errorMessage(l.message),
        (r) => successMessage('Удаление продукта', 'успешно произведено'));

    loadData();
  }

  createProduct() async {
    await Get.toNamed(Routes.ADD_PRODUCT);
    loadData();
  }

  openNotification() async {
    await Get.toNamed(Routes.NOTIFICATION);
    loadData();
  }
}
