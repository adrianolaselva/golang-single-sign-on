import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import {LoginComponent} from "./login/login.component";
import {OauthScopeComponent} from "./oauth-scope/oauth-scope.component";
import {DashboardHomeComponent } from "./dashboard-home/dashboard-home.component";
import {CallbackComponent} from "./callback/callback.component";
import {AuthGuardService} from "./auth/auth-guard.service";
import {LogOffComponent} from "./log-off/log-off.component";
import {ClientsComponent} from "./clients/clients.component";
import {UsersComponent} from "./users/users.component";
import {RolesComponent} from "./roles/roles.component";


const routes: Routes = [
  {path: 'auth/login', component: LoginComponent},
  {path: 'auth/callback', component: CallbackComponent},
  {path: 'auth/accepted-scopes', component: OauthScopeComponent},
  {
    path: 'admin',
    component: DashboardHomeComponent,
    canActivate: [AuthGuardService]
  },
  {
    path: 'admin/users',
    component: UsersComponent,
    canActivate: [AuthGuardService]
  },
  {
    path: 'admin/clients',
    component: ClientsComponent,
    canActivate: [AuthGuardService]
  },
  {
    path: 'admin/roles',
    component: RolesComponent,
    canActivate: [AuthGuardService]
  },
  {
    path: '',
    component: DashboardHomeComponent,
    canActivate: [AuthGuardService]
  },
  {path: 'logoff', component: LogOffComponent},
  {path: '**', redirectTo: 'admin' }
];

@NgModule({
  imports: [RouterModule.forRoot(routes, {useHash: true})],
  exports: [RouterModule]
})
export class AppRoutingModule { }
