import 'package:demo_store/internal/core/constant/formatter.dart';
import 'package:demo_store/internal/entity/notification/notification.dart';
import 'package:demo_store/widgets/card.dart';
import 'package:flutter/material.dart';

import 'package:get/get.dart';

import '../controllers/notification_controller.dart';

class NotificationView extends GetView<NotificationController> {
  const NotificationView({Key? key}) : super(key: key);
  @override
  Widget build(BuildContext context) {
    return Scaffold(
        appBar: AppBar(
          title: const Text('Уведомления'),
          centerTitle: true,
        ),
        body: c.obx((state) => _notificationList(state),
            onEmpty: Container(
              padding: const EdgeInsets.only(top: 50),
              child: const Row(
                mainAxisAlignment: MainAxisAlignment.center,
                children: <Widget>[
                  Icon(
                    Icons.info_outline,
                    color: Colors.grey,
                  ),
                  SizedBox(
                    width: 10,
                  ),
                  Text("Уведомлений нет",
                      style: TextStyle(color: Colors.grey, fontSize: 16))
                ],
              ),
            )));
  }

  Widget _notificationList(NotificationList notificationList) {
    final data = notificationList.messageList!;

    return ListView.builder(
        itemCount: data.length,
        itemBuilder: (context, index) {
          final row = data[index];
          return Card(
            child: Padding(
              padding: const EdgeInsets.all(8.0),
              child: Column(
                children: [
                  Padding(
                    padding: const EdgeInsets.only(top: 10, left: 10),
                    child: Row(
                      children: <Widget>[
                        Row(
                          children: <Widget>[
                            Text(
                              row.title,
                              style:
                                  const TextStyle(fontWeight: FontWeight.bold),
                            )
                          ],
                        ),
                      ],
                    ),
                  ),
                  Padding(
                    padding: const EdgeInsets.only(top: 5, left: 10),
                    child: Row(
                      children: <Widget>[
                        Row(
                          children: <Widget>[
                            Text(
                              Formatter.dft.format(
                                  DateTime.parse(row.createdDate).toLocal()),
                              style:
                                  const TextStyle(fontWeight: FontWeight.w400),
                            )
                          ],
                        ),
                      ],
                    ),
                  ),
                  Padding(
                    padding: const EdgeInsets.symmetric(
                        horizontal: 10, vertical: 15),
                    child: Row(
                      children: <Widget>[
                        Text(
                          row.message,
                        ),
                      ],
                    ),
                  ),
                ],
              ),
            ),
          );
        });
  }

  NotificationController get c => controller;
}
