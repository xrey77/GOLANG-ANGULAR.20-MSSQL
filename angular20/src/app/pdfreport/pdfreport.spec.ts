import { ComponentFixture, TestBed } from '@angular/core/testing';

import { Pdfreport } from './pdfreport';

describe('Pdfreport', () => {
  let component: Pdfreport;
  let fixture: ComponentFixture<Pdfreport>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [Pdfreport]
    })
    .compileComponents();

    fixture = TestBed.createComponent(Pdfreport);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
