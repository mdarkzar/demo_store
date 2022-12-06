import 'package:flutter/material.dart';
import 'package:flutter/rendering.dart';
import 'dart:math';

final textColor1 = Color(0xff516190);
final grey600 = Colors.grey[600];
final black2 = Color(0xff353a3d);
final progressBar1 = Color(0xff17a2b8);
final whitebackground = Color(0xfff8f8f8);

const darkBackground = Color(0xff1e232d);

const mainThemeColor = Color(0xff264653);
const mainBackgroundColor = Color(0xfff4f1de);

class HexColor extends Color {
  static int _getColorFromHex(String hexColor) {
    hexColor = hexColor.toUpperCase().replaceAll("#", "");
    if (hexColor.length == 6) {
      hexColor = "FF" + hexColor;
    }
    return int.parse(hexColor, radix: 16);
  }

  HexColor(final String hexColor) : super(_getColorFromHex(hexColor));
}

ThemeData createTheme() {
  final theme = ThemeData(
      primarySwatch: generateMaterialColor(mainThemeColor),
      primaryColor: mainThemeColor,
      scaffoldBackgroundColor: mainBackgroundColor);

  theme.copyWith(
    appBarTheme: const AppBarTheme(
      color: mainThemeColor,
    ),
    colorScheme: theme.colorScheme.copyWith(
      primary: mainThemeColor,
      secondary: mainThemeColor,
    ),
  );

  return theme;
}

MaterialColor generateMaterialColor(Color color) {
  return MaterialColor(color.value, {
    50: tintColor(color, 0.9),
    100: tintColor(color, 0.8),
    200: tintColor(color, 0.6),
    300: tintColor(color, 0.4),
    400: tintColor(color, 0.2),
    500: color,
    600: shadeColor(color, 0.1),
    700: shadeColor(color, 0.2),
    800: shadeColor(color, 0.3),
    900: shadeColor(color, 0.4),
  });
}

int tintValue(int value, double factor) =>
    max(0, min((value + ((255 - value) * factor)).round(), 255));

Color tintColor(Color color, double factor) => Color.fromRGBO(
    tintValue(color.red, factor),
    tintValue(color.green, factor),
    tintValue(color.blue, factor),
    1);

int shadeValue(int value, double factor) =>
    max(0, min(value - (value * factor).round(), 255));

Color shadeColor(Color color, double factor) => Color.fromRGBO(
    shadeValue(color.red, factor),
    shadeValue(color.green, factor),
    shadeValue(color.blue, factor),
    1);
