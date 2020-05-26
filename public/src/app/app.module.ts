import { BrowserModule } from '@angular/platform-browser';
import { NgModule, LOCALE_ID } from '@angular/core';
import {FormsModule, ReactiveFormsModule} from '@angular/forms';
import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { LoginComponent } from './login/login.component';
import {SharedModule} from "./shared/shared.module";
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { OauthScopeComponent } from './oauth-scope/oauth-scope.component';
import {HttpClientModule} from "@angular/common/http";
import {CoreModule} from "./core/core.module";
import { DashboardHomeComponent } from './dashboard-home/dashboard-home.component';
import { CallbackComponent } from './callback/callback.component';
import {AuthService} from "./auth/auth.service";
import {AuthGuardService} from "./auth/auth-guard.service";
import {JWT_OPTIONS, JwtModule} from "@auth0/angular-jwt";
import { LogOffComponent } from './log-off/log-off.component';
import { ClientsComponent } from './clients/clients.component';
import { RolesComponent } from './roles/roles.component';
import { UsersComponent } from './users/users.component';
import {InterceptorModule} from "./interceptors/interceptor.module";
import {AjaxLoadBusyDirective} from "./directive/ajax-load.directive";


@NgModule({
  declarations: [
    AppComponent,
    LoginComponent,
    OauthScopeComponent,
    DashboardHomeComponent,
    CallbackComponent,
    LogOffComponent,
    ClientsComponent,
    RolesComponent,
    UsersComponent,
    AjaxLoadBusyDirective
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    FormsModule,
    ReactiveFormsModule,
    SharedModule,
    BrowserAnimationsModule,
    HttpClientModule,
    CoreModule,
    InterceptorModule,
    JwtModule.forRoot({
      jwtOptionsProvider: {
        provide: JWT_OPTIONS,
        useFactory: jwtOptionsFactory
      },
      config: {
        // whitelistedDomains: ["localhost"],
        headerName: "Authorization",
        authScheme: "Bearer"
      }
    })
  ],
  providers: [
    AuthGuardService,
    AuthService,
    {provide: LOCALE_ID, useValue: 'pt-BR'}
  ],
  bootstrap: [AppComponent]
})
export class AppModule { }

export function jwtOptionsFactory(tokenService) {
  return {
    tokenGetter: () => {
      return sessionStorage.getItem("access_token");
    }
  }
}
