export interface Product {
  id: number;
  name: string;
  price: number;
  creator_id: number;
  created_date: Date;
  storage: string;
}

export interface Storage {
  id: number;
  name: string;
}
