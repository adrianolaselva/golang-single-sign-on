import { Component, OnInit, Input, ContentChild, AfterContentInit } from '@angular/core';
import { NgModel, FormControlName } from '@angular/forms';

@Component({
  selector: 'app-dashboard-content-component',
  templateUrl: './dashboard-content.component.html'
})
export class DashboardContentComponent implements OnInit, AfterContentInit {

  @Input() breadcrumbs: string;
  @Input() title: string;
  @Input() subtitle: boolean;
  @Input() loadAnimationEnable: boolean = false;

  input: any;

  @ContentChild(NgModel) model: NgModel;
  @ContentChild(FormControlName) control: FormControlName;

  constructor() { }

  ngOnInit(): void {
  }

  ngAfterContentInit(): void {
    this.input = this.model || this.control;
  }
}
