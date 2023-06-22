import 'package:demo_store/internal/repository/product.dart';
import 'package:demo_store/widgets/notification.dart';
import 'package:flutter/widgets.dart';
import 'package:get/get.dart';

class AddProductController extends GetxController with StateMixin {
  final ProductRepository productRepo;
  AddProductController(this.productRepo);

  final nameController = TextEditingController();
  final priceController = TextEditingController();
  final stID = RxInt(1);
  final formKey = GlobalKey<FormState>();

  @override
  void onInit() {
    super.onInit();
    change(null, status: RxStatus.success());
  }

  addProduct() async {
    if (!formKey.currentState!.validate()) {
      return;
    }

    change(null, status: RxStatus.loading());

    final r = await productRepo.create(nameController.value.text,
        double.parse(priceController.value.text), stID.value);
    r.fold((l) => errorMessage(l.message), (r) {
      Get.back();
      successMessage('Создание продукта', 'успешно произведено');
    });
  }

  selectStorage(int? value) {
    stID.value = value!;
  }
}
