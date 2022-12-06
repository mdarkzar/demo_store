import 'package:demo_store/resource/colors.dart';
import 'package:flutter/material.dart';
import 'package:get/get.dart';

errorMessage(String message,
    {Duration duration = const Duration(milliseconds: 1500)}) {
  Get.snackbar('Ошибка', message,
      backgroundColor: darkBackground,
      colorText: Colors.grey[300],
      leftBarIndicatorColor: Colors.red,
      borderRadius: 0,
      animationDuration: const Duration(milliseconds: 500),
      duration: duration,
      margin: const EdgeInsets.all(0));
}

successMessage(String title, message,
    {Duration duration = const Duration(milliseconds: 1500)}) {
  Get.snackbar(title, message,
      backgroundColor: darkBackground,
      colorText: Colors.white,
      leftBarIndicatorColor: Colors.green,
      borderRadius: 0,
      animationDuration: const Duration(milliseconds: 500),
      duration: duration,
      margin: const EdgeInsets.all(0));
}

warnMessage(String message,
    {Duration duration = const Duration(milliseconds: 1500)}) {
  Get.snackbar('Внимание', message,
      backgroundColor: darkBackground,
      colorText: Colors.black,
      leftBarIndicatorColor: Colors.yellow,
      borderRadius: 0,
      animationDuration: const Duration(milliseconds: 500),
      duration: duration,
      margin: const EdgeInsets.all(0));
}
