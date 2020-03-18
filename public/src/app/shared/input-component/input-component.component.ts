import { Component, OnInit, Input, ContentChild, AfterContentInit } from '@angular/core';
import { NgModel, FormControlName } from '@angular/forms';

@Component({
  selector: 'app-input-text-component',
  templateUrl: './input-component.component.html'
})
export class InputComponentComponent implements OnInit, AfterContentInit {

  @Input() label: string;
  @Input() errorMessage: string;
  @Input() enableSuccess: boolean;

  input: any;

  @ContentChild(NgModel) model: NgModel;
  @ContentChild(FormControlName) control: FormControlName;

  constructor() { }

  ngOnInit(): void {
  }

  ngAfterContentInit(): void {
    this.input = this.model || this.control;

    if(this.input === undefined) {
      throw new Error("Este componente não pode ser usado com a diretiva NgModel");
    }
  }

  hasSuccess(): boolean {
    return this.input.valid && (this.input.dirty || this.input.touched);
  }

  hasError(): boolean {
    return this.input.invalid && (this.input.dirty || this.input.touched);
  }

}
