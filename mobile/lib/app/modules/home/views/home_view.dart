import 'package:demo_store/app/modules/home/views/drawer.dart';
import 'package:demo_store/internal/core/constant/formatter.dart';
import 'package:demo_store/internal/entity/product/product.dart';
import 'package:demo_store/widgets/buttons.dart';
import 'package:demo_store/widgets/card.dart';
import 'package:flutter/material.dart';

import 'package:get/get.dart';
import 'package:material_design_icons_flutter/material_design_icons_flutter.dart';

import '../controllers/home_controller.dart';

class HomeView extends GetView<HomeController> {
  const HomeView({Key? key}) : super(key: key);
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: const Text('Demo Store'),
        actions: [
          IconButton(
              onPressed: c.openNotification, icon: const Icon(MdiIcons.bell)),
          IconButton(
              onPressed: c.createProduct, icon: const Icon(MdiIcons.plus))
        ],
      ),
      drawer: const AppDrawer(),
      body: c.obx((state) => _productList(state),
          onEmpty: const Center(
            child: Text('Продуктов нет'),
          )),
    );
  }

  _productList(ProductList productList) {
    final data = productList.productList!;

    return RefreshIndicator(
        onRefresh: () => controller.loadData(),
        child: ListView.builder(
            itemCount: data.length,
            itemBuilder: (context, index) {
              final row = data[index];
              return Padding(
                padding: const EdgeInsets.all(3.0),
                child: Card(
                  child: Padding(
                    padding: const EdgeInsets.all(10.0),
                    child: Column(
                      children: [
                        Row(
                          children: [
                            Flexible(
                                child: Text(
                              row.name,
                              style:
                                  const TextStyle(fontWeight: FontWeight.w600),
                            ))
                          ],
                        ),
                        const SizedBox(
                          height: 15,
                        ),
                        cardField(
                          'Стоимость',
                          Formatter.formatCurrency.format(row.price),
                        ),
                        const SizedBox(
                          height: 15,
                        ),
                        Row(
                          mainAxisAlignment: MainAxisAlignment.end,
                          children: [
                            standartButton('Удалить', () {
                              Get.defaultDialog(
                                  title: 'Удаление продукта',
                                  middleText:
                                      'Вы уверены, что хотите удалить продукт?',
                                  actions: [
                                    standartButton('Подтверждаю', () async {
                                      Get.back();
                                      c.deleteProduct(row.id);
                                    }, const Size(100, 40), color: Colors.red),
                                    standartButton('Отмена', () {
                                      Get.back();
                                    }, const Size(100, 40)),
                                  ]);
                            }, const Size(80, 35), color: Colors.red),
                          ],
                        )
                      ],
                    ),
                  ),
                ),
              );
            }));
  }

  HomeController get c => controller;
}
