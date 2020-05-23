import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';
import {FormsModule, ReactiveFormsModule} from '@angular/forms';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { LoginComponent } from './login/login.component';
import {SharedModule} from "./shared/shared.module";
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { OauthScopeComponent } from './oauth-scope/oauth-scope.component';
import {HttpClientModule} from "@angular/common/http";
import {CoreModule} from "./core/core.module";
import { DashboardHomeComponent } from './dashboard/dashboard-home.component';
import { CallbackComponent } from './callback/callback.component';
import {AuthService} from "./auth/auth.service";
import {AuthGuardService} from "./auth/auth-guard.service";
import {JWT_OPTIONS, JwtModule} from "@auth0/angular-jwt";
import { LogOffComponent } from './log-off/log-off.component';

@NgModule({
  declarations: [
    AppComponent,
    LoginComponent,
    OauthScopeComponent,
    DashboardHomeComponent,
    CallbackComponent,
    LogOffComponent
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
    JwtModule.forRoot({
      jwtOptionsProvider: {
        provide: JWT_OPTIONS,
        useFactory: jwtOptionsFactory
      },
      config: {
        whitelistedDomains: ["localhost"],
        headerName: "Authorization",
        authScheme: "Bearer"
      }
    })
  ],
  providers: [
    AuthGuardService,
    AuthService
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
