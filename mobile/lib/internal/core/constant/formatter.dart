import 'package:flutter/services.dart';
import 'package:intl/intl.dart';
import 'package:mask_text_input_formatter/mask_text_input_formatter.dart';

class Formatter {
  static NumberFormat formatCurrency = NumberFormat("#,##0", "ru_RU");
  static MaskTextInputFormatter phoneFormatter = MaskTextInputFormatter(
      mask: '+998-##-###-##-##', filter: {"#": RegExp(r'[0-9]')});
  static DateFormat dft = DateFormat('dd MMM yyyy HH:mm', "ru_RU");
  static DateFormat df = DateFormat('d MMM yyyy', "ru_RU");
  static MaskTextInputFormatter macFormatter = MaskTextInputFormatter(
      mask: '##:##:##:##:##:##', filter: {"#": RegExp(r'[0-9-A-Z-a-z]')});

  /// A [TextInputFormatter] that takes in digits `[0-9]` only.
  static final TextInputFormatter doubleFormatter =
      FilteringTextInputFormatter.allow(RegExp(r'[0-9\.]'));
}
