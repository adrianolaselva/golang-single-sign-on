import { Component, OnInit } from '@angular/core';
import {FormBuilder, FormControl, FormGroup, Validators} from "@angular/forms";
import {ActivatedRoute} from "@angular/router";
import {LoginService} from "./login.service";
import {LoginRequestModel} from "./login.request.model";

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent implements OnInit {

  formLogin;

  responseType: string;
  clientId: string;
  redirectUri: string;
  state: string;
  scope: string;

  constructor(private formBuilder: FormBuilder, private route: ActivatedRoute, private loginService: LoginService) { }

  ngOnInit(): void {
    this.route.queryParams.subscribe(params => {
      this.responseType = params['response_type'];
      this.clientId = params['client_id'];
      this.redirectUri = params['redirect_uri'];
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

    this.loginService.login(loginRequest).subscribe(data => {
      window.location.href = `${data.redirect_uri}?token_type=${data.access_token.token_type}&expires_in=${data.access_token.expires_in}&access_token=${data.access_token.access_token}&refresh_token=${data.access_token.refresh_token}&state=${data.access_token.state}`
    }, error => {
      console.log(error.error.error);
    })
  }

  get name() {
    return this.formLogin.username;
  }

  get password() {
    return this.formLogin.password
  }

}
