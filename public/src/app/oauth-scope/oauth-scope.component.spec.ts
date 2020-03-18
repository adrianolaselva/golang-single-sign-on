import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { OauthScopeComponent } from './oauth-scope.component';

describe('OauthScopeComponent', () => {
  let component: OauthScopeComponent;
  let fixture: ComponentFixture<OauthScopeComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ OauthScopeComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(OauthScopeComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
