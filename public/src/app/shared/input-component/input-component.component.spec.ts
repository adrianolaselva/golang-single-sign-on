import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { InputComponentComponent } from './input-component.component';

describe('InputComponent', () => {
  let component: InputComponentComponent;
  let fixture: ComponentFixture<InputComponentComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ InputComponentComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(InputComponentComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
