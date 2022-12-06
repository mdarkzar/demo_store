import 'package:demo_store/app/routes/app_pages.dart';
import 'package:demo_store/internal/repository/user.dart';
import 'package:get/get.dart';

class RootController extends GetxController with StateMixin {
  final UserRepository _userRepository;
  RootController(this._userRepository);

  @override
  void onInit() {
    super.onInit();
  }

  @override
  void onReady() {
    checkAuth();
    super.onReady();
  }

  checkAuth() async {
    final r = await _userRepository.profile();
    r.fold((l) {
      Get.offNamed(Routes.LOGIN);
    }, (r) {
      Get.put(r, permanent: true);
      Get.offNamed(Routes.HOME);
    });
  }
}
