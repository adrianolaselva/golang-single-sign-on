import { Component, OnInit } from '@angular/core';
import {FormBuilder, FormControl, FormGroup, Validators} from "@angular/forms";

@Component({
  selector: 'app-oauth-scope',
  templateUrl: './oauth-scope.component.html',
  styleUrls: ['./oauth-scope.component.css']
})
export class OauthScopeComponent implements OnInit {

  formAuthScopes;

  scopesList: any[] = ['Edit users','Edit rules','Edit applications'];

  constructor(private formBuilder: FormBuilder) { }

  ngOnInit(): void {
    this.formAuthScopes = this.formBuilder.group({

    });
  }

  onSubmit(data) {

  }
}
