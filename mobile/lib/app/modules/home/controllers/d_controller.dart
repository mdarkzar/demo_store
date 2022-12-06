import 'package:demo_store/app/routes/app_pages.dart';
import 'package:demo_store/internal/repository/product.dart';
import 'package:demo_store/internal/repository/user.dart';
import 'package:demo_store/widgets/notification.dart';
import 'package:get/get.dart';

class DController extends GetxController with StateMixin {
  final UserRepository userRepo;
  DController(this.userRepo);

  @override
  void onInit() {
    super.onInit();
    loadData();
  }

  loadData() async {
    change(null, status: RxStatus.loading());
    final r = await userRepo.profile();
    r.fold((l) => errorMessage(l.message),
        (r) => change(r, status: RxStatus.success()));
  }

  logout() async {
    change(null, status: RxStatus.loading());
    final r = await userRepo.logout();
    r.fold((l) => errorMessage(l.message), (r) => Get.offNamed(Routes.LOGIN));
  }
}
