import { Component, OnInit } from '@angular/core';
import {FormBuilder, FormControl, FormGroup, Validators} from "@angular/forms";

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent implements OnInit {

  formLogin;

  constructor(private formBuilder: FormBuilder) { }

  ngOnInit(): void {
    this.formLogin = this.formBuilder.group({
      username: '',
      password: ''
    })

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
    console.log(loginData)
    this.formLogin.reset()
  }

  get name() {
    return this.formLogin.username;
  }

  get password() {
    return this.formLogin.password
  }

}
