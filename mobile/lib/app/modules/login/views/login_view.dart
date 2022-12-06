import 'package:demo_store/widgets/buttons.dart';
import 'package:demo_store/widgets/padding.dart';
import 'package:demo_store/widgets/textfield.dart';
import 'package:flutter/material.dart';

import 'package:get/get.dart';
import 'package:material_design_icons_flutter/material_design_icons_flutter.dart';
import 'package:sizer/sizer.dart';

import '../controllers/login_controller.dart';

class LoginView extends GetView<LoginController> {
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: controller.obx(
        (state) => _form(),
      ),
    );
  }

  Widget _form() {
    return centerSingleChildScroll(
        width: 70.w,
        child: Form(
            key: c.formKey,
            child: Obx(
              () => Column(children: [
                pB(_logo(), 20),
                pB(_loginField(), 10),
                pB(_passwordField(), 10),
                pB(_obscureButton(), 15),
                pB(_loginButton(), 15),
              ]),
            )));
  }

  Widget _logo() {
    return Image.asset(
      'assets/images/logo.png',
      width: 50.0.w,
    );
  }

  Widget _loginField() {
    return standartTextField(
        controller: c.loginController,
        autofillHints: [AutofillHints.email],
        hintText: "Логин",
        validator: (value) {
          if (value!.isEmpty) return "Введите логин";
          return null;
        },
        icon: MdiIcons.email);
  }

  Widget _passwordField() {
    return standartTextField(
        controller: c.passwordController,
        autofillHints: [AutofillHints.password],
        hintText: "Пароль",
        obscureText: controller.obscureText.value,
        validator: (value) {
          if (value!.isEmpty) return "Введите пароль";
          return null;
        },
        icon: MdiIcons.key);
  }

  Widget _obscureButton() {
    return TextButton(
        onPressed: controller.toggleObscure,
        child: Icon(
          controller.obscureText.value ? MdiIcons.eyeOff : MdiIcons.eye,
          color: Colors.grey[700],
        ));
  }

  Widget _loginButton() {
    return standartButton("Войти", () async {
      if (!c.formKey.currentState!.validate()) {
        return;
      }

      await controller.auth();

      c.passwordController.clear(); // очистить пароль
    }, const Size(200, 44));
  }

  LoginController get c => controller;
}
