import 'package:dartz/dartz.dart';
import 'package:demo_store/internal/core/error/failures.dart';
import 'package:demo_store/internal/datasources/api.dart';
import 'package:demo_store/internal/entity/product/product.dart';

class ProductRepository {
  final DataSource source;

  ProductRepository(this.source);

  Future<Either<Failure, ProductList>> loadAll() async {
    final res = await source.get("/product/load");

    return res.fold((l) => Left(l), (r) => Right(ProductList.fromJson(r)));
  }

  Future<Either<Failure, void>> create(String name, double price) async {
    final res =
        await source.post("/product/create", {"name": name, "price": price});

    return res.fold((l) => Left(l), (r) => const Right(null));
  }

  Future<Either<Failure, void>> delete(int productID) async {
    final res = await source.post("/product/remove", {"product_id": productID});

    return res.fold((l) => Left(l), (r) => const Right(null));
  }
}
