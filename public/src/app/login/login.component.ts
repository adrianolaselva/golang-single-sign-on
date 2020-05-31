import { Component, OnInit } from '@angular/core';
import {FormBuilder, FormControl, FormGroup, Validators} from "@angular/forms";
import {ActivatedRoute, Router} from "@angular/router";
import {LoginService} from "./login.service";
import {LoginRequestModel} from "./login.request.model";
import {environment} from "../../environments/environment";
import {JwtHelperService} from "@auth0/angular-jwt";
import {ToastrService} from "ngx-toastr";

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent implements OnInit {

  // loadAnimationEnable: boolean = true;
  formLogin;
  responseType: string;
  clientId: string;
  redirectUri: string;
  state: string;
  scope: string;
  errorMessage: string = null;
  successMessage: string = null;

  constructor(private formBuilder: FormBuilder,
              private route: ActivatedRoute,
              private loginService: LoginService,
              public jwtHelper: JwtHelperService,
              public router: Router,
              private toastr: ToastrService) { }

  ngOnInit(): void {
    if (!this.jwtHelper.isTokenExpired(sessionStorage.getItem('access_token'))) {
      this.router.navigate(['admin']);
      return;
    }

    this.route.queryParams.subscribe(params => {
      this.responseType = params['response_type'] ? params['response_type'] : "token";
      this.clientId = params['client_id'] ? params['client_id'] : environment.CLIENT_ID;
      this.redirectUri = params['redirect_uri'] ? params['redirect_uri'] : environment.REDIRECT_URL;
      this.state = params['state'];
      this.scope = params['scope'];
    });

    this.formLogin = this.formBuilder.group({
      username: '',
      password: ''
    });

    this.formLogin = new FormGroup({
      username: new FormControl(this.formLogin.username, [
        Validators.required,
        Validators.minLength(3),
        Validators.maxLength(90),
      ]),
      password: new FormControl(this.formLogin.password, [
        Validators.required,
        Validators.minLength(3),
        Validators.maxLength(90),
      ]),
    })
  }

  onSubmit(loginData) {
    let loginRequest = new LoginRequestModel();
    loginRequest.client_id = this.clientId;
    loginRequest.response_type = this.responseType;
    loginRequest.redirect_uri = this.redirectUri;
    loginRequest.state = this.state;
    loginRequest.scope = this.scope;
    loginRequest.username = loginData.username;
    loginRequest.password = loginData.password;

    // this.loadAnimationEnable = true;
    this.loginService.login(loginRequest).pipe(
      // finalize(() => this.loadAnimationEnable = false)
    ).subscribe(data => {
      this.errorMessage = null;
      // this.successMessage = "autenticado com sucesso";
      this.toastr.success("autenticado com sucesso", "sucesso");

      if(data.response_type == "token") {
        window.location.href = `${data.redirect_uri}?token_type=${data.access_token.token_type}&expires_in=${data.access_token.expires_in}&access_token=${data.access_token.access_token}&refresh_token=${data.access_token.refresh_token}&state=${data.access_token.state}`
      }

      if(data.response_type == "code") {
        window.location.href = `${data.redirect_uri}?state=${data.state}&code=${data.code}`
      }
    }, error => {
      this.successMessage = null;
      // this.errorMessage = error.error.error;
      if (error.error.error) {
        this.toastr.error(error.error.error, "falha");
      }else {
        this.toastr.error(error.message, "falha");
      }
    })
  }

  get name() {
    return this.formLogin.username;
  }

  get password() {
    return this.formLogin.password
  }

}
