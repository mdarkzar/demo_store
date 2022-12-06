import 'package:flutter/material.dart';

import '../resource/colors.dart';

Widget standartButton(String? text, dynamic onPressed, Size size,
    {double fontsize = 16,
    Icon? icon,
    Color? color = mainThemeColor,
    bool rounded = false}) {
  final ButtonStyle raisedButtonStyle = ElevatedButton.styleFrom(
    onPrimary: Colors.white12,
    primary: color,
    minimumSize: size,
    padding: rounded
        ? const EdgeInsets.symmetric(horizontal: 18, vertical: 9)
        : const EdgeInsets.symmetric(horizontal: 16),
    shape: rounded
        ? const StadiumBorder()
        : const RoundedRectangleBorder(
            borderRadius: BorderRadius.all(Radius.circular(8)),
          ),
  );

  return ElevatedButton(
    onPressed: onPressed,
    child: Row(
      mainAxisAlignment: MainAxisAlignment.center,
      children: <Widget>[
        icon ?? const SizedBox(),
        SizedBox(
          width: text != null && text != '' ? 7 : 0,
        ),
        text != null && text != ''
            ? Text(
                text,
                style: TextStyle(fontSize: fontsize, color: Colors.white),
              )
            : const SizedBox()
      ],
    ),
    style: raisedButtonStyle,
  );
}
