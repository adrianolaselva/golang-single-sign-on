import {Directive, ElementRef, Input, OnInit, Renderer2} from "@angular/core";
import {interval, Subscription} from "rxjs";
import {AjaxLoadNotifierService} from "./ajax-load-notifier.service";

@Directive({
  selector: '[ajax-load-derective]'
})
export class AjaxLoadBusyDirective implements OnInit {

  @Input() showDelay: number = 50;
  @Input() hideDelay: number = 1000;
  hideTimer: Subscription;
  showTimer: Subscription;

  constructor(private el: ElementRef, private renderer: Renderer2, private alns: AjaxLoadNotifierService) {
  }

  ngOnInit(): void {
    this.alns.busy.subscribe(busy => {
      if (busy) {
        this.cancelPendingHide();
        if (!this.showTimer) {
          this.showTimer = interval(this.showDelay).subscribe(() => {
            this.renderer.setStyle(this.el.nativeElement, 'display', 'block');
            this.showTimer.unsubscribe();
            this.showTimer = null;
          });
        }
      }
    });

    this.alns.busy.subscribe(busy => {
      if (!busy) {
        this.cancelPendingShow();
        if (!this.hideTimer) {
          this.hideTimer = interval(this.hideDelay).subscribe(() => {
            this.renderer.setStyle(this.el.nativeElement, 'display', 'none');
            this.hideTimer.unsubscribe(); this.hideTimer = null;
          });
        }
      }
    });
  }

  cancelPendingHide() {
    if (this.hideTimer) {
      this.hideTimer.unsubscribe();
      delete this.hideTimer;
    }
  }

  cancelPendingShow() {
    if (this.showTimer) {
      this.showTimer.unsubscribe();
      delete this.showTimer;
    }
  }
}
