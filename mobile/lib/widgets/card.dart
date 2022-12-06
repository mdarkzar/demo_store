import 'package:flutter/widgets.dart';

final roundedRectangleBorder = RoundedRectangleBorder(
  borderRadius: BorderRadius.circular(15.0),
);

cardField(String key, value, {Color? textColor}) {
  return Padding(
    padding: const EdgeInsets.only(bottom: 5),
    child: Row(
      children: [
        Text(
          '$key:',
          style: const TextStyle(fontWeight: FontWeight.w600),
        ),
        const SizedBox(
          width: 5,
        ),
        Flexible(
          child: Text(
            value,
            style: TextStyle(color: textColor),
          ),
        )
      ],
    ),
  );
}
