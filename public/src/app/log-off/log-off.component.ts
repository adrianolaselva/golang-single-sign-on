import { Component, OnInit } from '@angular/core';
import {AuthService} from "../auth/auth.service";
import {Router} from "@angular/router";

@Component({
  selector: 'app-log-off',
  templateUrl: './log-off.component.html'
})
export class LogOffComponent implements OnInit {

  constructor(public router: Router) { }

  ngOnInit(): void {
    sessionStorage.clear();
    this.router.navigate(['auth/login']);
  }

}
