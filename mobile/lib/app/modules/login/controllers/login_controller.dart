import 'package:demo_store/app/routes/app_pages.dart';
import 'package:demo_store/internal/repository/user.dart';
import 'package:demo_store/widgets/notification.dart';
import 'package:flutter/widgets.dart';
import 'package:get/get.dart';

class LoginController extends GetxController with StateMixin {
  final UserRepository userRepo;
  LoginController(this.userRepo);

  final formKey = GlobalKey<FormState>();
  final loginController = TextEditingController();
  final passwordController = TextEditingController();

  @override
  void onInit() {
    change(null, status: RxStatus.success());
    super.onInit();
  }

  var obscureText = true.obs;

  toggleObscure() {
    obscureText.value = !obscureText.value;
  }

  auth() async {
    change(null, status: RxStatus.loading());
    final resultOrFail = await userRepo.login(
        loginController.text.trim(), passwordController.text.trim());

    resultOrFail.fold((l) {
      errorMessage(l.message);
    }, (r) async {
      Get.offNamed(Routes.HOME);
    });

    change(null, status: RxStatus.success());
  }
}
