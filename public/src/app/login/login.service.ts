import { Injectable } from '@angular/core';
import {HttpClient} from "@angular/common/http";
import {LoginRequestModel} from "./login.request.model";
import {Observable} from "rxjs";
import {LoginResponseModel} from "./login.response.model";


@Injectable()
export class LoginService {

  constructor(
    private http: HttpClient
  ) {}

  login(loginRequest: LoginRequestModel): Observable<LoginResponseModel> {
    return this.http.post<LoginResponseModel>(
      `http://localhost:9091/oauth2/login`, JSON.stringify(loginRequest), {
      headers: {
        "Content-Type": "application/json"
      }
    })
  }

}
