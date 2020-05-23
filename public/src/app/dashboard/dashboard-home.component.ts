import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-dashboard-home',
  templateUrl: './dashboard-home.component.html',
  styleUrls: ['./dashboard-home.component.css']
})
export class DashboardHomeComponent implements OnInit {

  constructor() { }

  ngOnInit(): void {
    console.log("dashboard")
    console.log(sessionStorage.getItem("access_token"))
  }

}
