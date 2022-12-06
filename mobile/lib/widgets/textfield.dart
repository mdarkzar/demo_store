import 'package:demo_store/resource/colors.dart';
import 'package:flutter/material.dart';
import 'package:flutter/services.dart';

Widget standartTextField(
    {required TextEditingController controller,
    String? Function(String?)? validator,
    String? hintText,
    IconData? icon,
    bool obscureText = false,
    bool disabled = false,
    TextInputFormatter? inputFormatters,
    bool requiredField = false,
    int? maxLines,
    Function(String)? onChanged,
    TextInputType? keyboardType,
    Iterable<String>? autofillHints}) {
  if (validator == null && requiredField) {
    validator = (String? value) {
      if (value == null) return "Введите $hintText";
      return null;
    };
  }

  return TextFormField(
    inputFormatters: inputFormatters != null ? [inputFormatters] : null,
    maxLines: obscureText ? 1 : maxLines,
    enabled: !disabled,
    style: const TextStyle(color: Colors.black),
    validator: validator,
    controller: controller,
    obscureText: obscureText,
    decoration: InputDecoration(
      hintText: hintText,
      hintStyle: TextStyle(color: grey600),
      prefixIcon: _ifHasIcon(icon),
      enabledBorder: UnderlineInputBorder(
        borderSide: BorderSide(color: grey600!),
      ),
    ),
    onChanged: onChanged,
    keyboardType: keyboardType,
    autofillHints: autofillHints,
  );
}

Widget? _ifHasIcon(IconData? icon) {
  return icon != null
      ? Icon(
          icon,
          color: grey600,
        )
      : null;
}

InputDecoration textDecoration(String hint) {
  return InputDecoration(
    alignLabelWithHint: true,
    labelText: hint,
    labelStyle: TextStyle(color: grey600),
    hintStyle: TextStyle(color: grey600),
    enabledBorder: UnderlineInputBorder(
      borderSide: BorderSide(color: grey600!),
    ),
  );
}
