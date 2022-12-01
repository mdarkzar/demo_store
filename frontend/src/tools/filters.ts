import moment from "moment";
moment.locale("ru");

export const numberFormatter = (nStr: any) => {
  if (!nStr && nStr !== 0) {
    return "";
  }
  nStr += "";
  const x = nStr.split(".");
  let x1 = x[0];
  const x2 = x.length > 1 ? "." + x[1] : "";
  const rgx = /(\d+)(\d{3})/;
  while (rgx.test(x1)) {
    x1 = x1.replace(rgx, "$1" + " " + "$2");
  }
  return x1 + x2;
};

export const dateFormat = (date: Date, format: string) => {
  if (!date) {
    return "";
  }
  return moment(date).format(format);
};
