import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-clients',
  templateUrl: './clients.component.html',
  styleUrls: ['./clients.component.css']
})
export class ClientsComponent implements OnInit {

  loadAnimationEnable: boolean = false;

  constructor() { }

  ngOnInit(): void {
  }

}
