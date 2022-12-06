import 'package:demo_store/internal/repository/notification.dart';
import 'package:demo_store/widgets/notification.dart';
import 'package:get/get.dart';

class NotificationController extends GetxController with StateMixin {
  final NotificationRepository productRepo;
  NotificationController(this.productRepo);

  @override
  void onInit() {
    super.onInit();
    loadData();
  }

  loadData() async {
    change(null, status: RxStatus.loading());
    final r = await productRepo.loadNew();
    r.fold(
        (l) => errorMessage(l.message),
        (r) => change(r,
            status: r.messageList == null || r.messageList!.isEmpty
                ? RxStatus.empty()
                : RxStatus.success()));
  }
}
