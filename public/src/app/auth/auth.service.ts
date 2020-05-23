
import { Injectable } from "@angular/core";
import { JwtHelperService } from "@auth0/angular-jwt";

@Injectable()
export class AuthService {

  constructor(public jwtHelper: JwtHelperService) {}

  public isAuthenticated(): boolean {
    const accessToken = sessionStorage.getItem('access_token');

    console.log(this.jwtHelper.decodeToken(accessToken));

    return !this.jwtHelper.isTokenExpired(accessToken);
  }
}
