import {NgModule} from '@angular/core';
import {CommonModule} from '@angular/common';
import {FormsModule, ReactiveFormsModule} from '@angular/forms';

import {InputComponentComponent} from "./input-component/input-component.component";
import {PasswordComponentComponent} from "./password-component/password-component.component";
import {DashboardContentComponent} from "./dashboard-content-component/dashboard-content.component";
import {AppRoutingModule} from "../app-routing.module";

@NgModule({
  declarations: [InputComponentComponent,PasswordComponentComponent, DashboardContentComponent],
  imports: [CommonModule, FormsModule, ReactiveFormsModule, AppRoutingModule],
  exports: [
    InputComponentComponent,
    PasswordComponentComponent,
    CommonModule,
    FormsModule,
    ReactiveFormsModule,
    DashboardContentComponent
  ]
})
export class SharedModule {

}
