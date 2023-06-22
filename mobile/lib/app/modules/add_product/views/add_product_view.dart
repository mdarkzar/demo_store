import 'package:demo_store/app/modules/add_product/controllers/storage_controller.dart';
import 'package:demo_store/widgets/buttons.dart';
import 'package:demo_store/widgets/padding.dart';
import 'package:demo_store/widgets/textfield.dart';
import 'package:flutter/material.dart';
import 'package:flutter/services.dart';

import 'package:get/get.dart';

import '../controllers/add_product_controller.dart';

class AddProductView extends GetView<AddProductController> {
  const AddProductView({super.key});

  @override
  Widget build(BuildContext context) {
    return Scaffold(
        appBar: AppBar(
          title: const Text('Добавление продукта'),
          centerTitle: true,
        ),
        body: c.obx((state) => _form()));
  }

  Widget _form() {
    return Form(
        key: c.formKey,
        child: Padding(
            padding: const EdgeInsets.all(16.0),
            child: Column(
              children: [
                standartTextField(
                    controller: c.nameController,
                    validator: (value) {
                      if (value!.isEmpty) return "Введите название";
                      return null;
                    },
                    hintText: "Название"),
                hSpace(1),
                standartTextField(
                    controller: c.priceController,
                    keyboardType: TextInputType.number,
                    inputFormatters: FilteringTextInputFormatter.digitsOnly,
                    validator: (value) {
                      if (value!.isEmpty) return "Введите стоимость";
                      return null;
                    },
                    hintText: "Стоимость"),
                hSpace(1),
                Container(
                  child: _storageList(),
                ),
                hSpace(4),
                saveButton(),
              ],
            )));
  }

  Widget _storageList() {
    return GetBuilder<StorageController>(
      builder: (controller) {
        return DropdownButtonFormField<int>(
          decoration: textDecoration('Склад'),
          value: c.stID.value,
          validator: (value) {
            if (value == null) return "Выберите склад";
            return null;
          },
          style: const TextStyle(fontSize: 16, color: Colors.black),
          isExpanded: true,
          items: controller.storageList != null
              ? controller.storageList!.value.map((row) {
                  return DropdownMenuItem<int>(
                    value: row.id,
                    child: Text(row.name),
                  );
                }).toList()
              : null,
          onChanged: c.selectStorage,
        );
      },
    );
  }

  Widget saveButton() {
    return standartButton("Сохранить", () async {
      await c.addProduct();
    }, const Size(100, 40));
  }

  AddProductController get c => controller;
}
