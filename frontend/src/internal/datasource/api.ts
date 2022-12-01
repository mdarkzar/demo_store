import { err, ok, Result } from "neverthrow";
import { Failure, ResponseFailure } from "@/internal/entity/failure";
import { KeyValue, Response } from "@/internal/entity/rest";
import axios from "axios";

export default class API {
  async post(url: string, params: KeyValue): Promise<Result<any, Failure>> {
    try {
      const { data } = await axios.post<Response>(`/api/v1${url}`, params);
      if (data.error) {
        return err(new ResponseFailure(data.error));
      } else if (data.result || data.result !== undefined) {
        return ok(data.result);
      } else {
        return err(new ResponseFailure("Сервер временно недоступен"));
      }
    } catch (errHttp) {
      return err(new ResponseFailure("Сервер временно недоступен"));
    }
  }
  async get(url: string): Promise<Result<any, Failure>> {
    try {
      const { data } = await axios.get<Response>(`/api/v1${url}`);

      if (data.error) {
        return err(new ResponseFailure(data.error));
      } else if (data.result || data.result !== undefined) {
        return ok(data.result);
      } else {
        return err(new ResponseFailure("Сервер временно недоступен"));
      }
    } catch (errHttp) {
      return err(new ResponseFailure("Сервер временно недоступен"));
    }
  }
}
