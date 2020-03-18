import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import {LoginComponent} from "./login/login.component";
import {OauthScopeComponent} from "./oauth-scope/oauth-scope.component";


const routes: Routes = [
  {path: 'auth/login', component: LoginComponent},
  {path: 'auth/accepted-scopes', component: OauthScopeComponent},
];

@NgModule({
  imports: [RouterModule.forRoot(routes, {useHash: true})],
  exports: [RouterModule]
})
export class AppRoutingModule { }
