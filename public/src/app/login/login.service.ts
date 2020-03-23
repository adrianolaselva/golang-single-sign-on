import { Injectable } from '@angular/core';
import {HttpClient} from "@angular/common/http";
import {LoginRequestModel} from "./login.request.model";
import {Observable} from "rxjs";
import {LoginResponseModel} from "./login.response.model";
import {environment} from "../../environments/environment";


@Injectable()
export class LoginService {

  host: string

  constructor(
    private http: HttpClient
  ) { }

  login(loginRequest: LoginRequestModel): Observable<LoginResponseModel> {
    return this.http.post<LoginResponseModel>(
      `${environment.MIDGARD_API}oauth2/login`, JSON.stringify(loginRequest), {
      headers: {
        "Content-Type": "application/json"
      }
    })
  }

}
