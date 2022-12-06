import 'package:demo_store/app/modules/home/controllers/d_controller.dart';
import 'package:flutter/material.dart';
import 'package:material_design_icons_flutter/material_design_icons_flutter.dart';
import 'package:get/get.dart';

class AppDrawer extends GetView<DController> {
  const AppDrawer({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Drawer(
      child: Container(
        color: Colors.white,
        child: ListView(
          // padding: EdgeInsets.fromLTRB(5, 150, 0, 30),
          children: <Widget>[
            DrawerHeader(
                child: controller.obx(
              (state) => ListTile(
                contentPadding: const EdgeInsets.only(top: 60),
                title: Text(
                  '${state.login}',
                  style: const TextStyle(fontSize: 18),
                ),
                dense: true,
                leading: const CircleAvatar(
                  backgroundColor: Colors.orange,
                  child: Text(
                    "D",
                    style: TextStyle(
                        color: Colors.white,
                        fontWeight: FontWeight.bold,
                        fontSize: 20),
                  ),
                ),
              ),
            )),
            ListTile(
              leading: Icon(
                MdiIcons.logout,
                size: 24,
                color: Colors.grey[600],
              ),
              title: Text(
                'Выйти из аккаунта',
                style: TextStyle(color: Colors.grey[600], fontSize: 16),
              ),
              onTap: () {
                // widget.userCubit.logout();
                controller.logout();
              },
            ),
            const SizedBox(
              height: 50,
            ),
          ],
        ),
      ),
    );
  }
}
