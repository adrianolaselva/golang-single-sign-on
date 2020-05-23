import { Component, OnInit } from '@angular/core';
import {ActivatedRoute, Router} from "@angular/router";

@Component({
  selector: 'app-callback',
  templateUrl: './callback.component.html'
})
export class CallbackComponent implements OnInit {

  constructor(private route: ActivatedRoute, private router: Router) { }

  ngOnInit(): void {
    this.route.queryParams.subscribe(params => {
      sessionStorage.setItem("token_type", params["token_type"]);
      sessionStorage.setItem("access_token", params["access_token"]);
      sessionStorage.setItem("expires_in", params["expires_in"]);
      sessionStorage.setItem("state", params["state"]);
    });
    this.router.navigate(["admin"]).then(r => console.log(r))
  }

}
