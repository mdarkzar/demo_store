import 'package:flutter/material.dart';

import 'package:get/get.dart';

import '../controllers/root_controller.dart';

class RootView extends GetView<RootController> {
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: const Text('Demo Store'),
        centerTitle: true,
      ),
      body: controller
          .obx((state) => const Center(child: CircularProgressIndicator())),
    );
  }
}
