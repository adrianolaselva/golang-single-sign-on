import {NgModule} from '@angular/core';
import {CommonModule} from '@angular/common';
import {FormsModule, ReactiveFormsModule} from '@angular/forms';

import {InputComponentComponent} from "./input-component/input-component.component";
import {PasswordComponentComponent} from "./password-component/password-component.component";

@NgModule({
  declarations: [InputComponentComponent,PasswordComponentComponent],
  imports: [CommonModule, FormsModule, ReactiveFormsModule],
  exports: [
    InputComponentComponent,
    PasswordComponentComponent,
    CommonModule,
    FormsModule,
    ReactiveFormsModule
  ]
})
export class SharedModule {

}
