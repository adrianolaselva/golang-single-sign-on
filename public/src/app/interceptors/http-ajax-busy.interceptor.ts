import { Injectable } from '@angular/core';
import {
  HttpRequest,
  HttpHandler,
  HttpEvent,
  HttpInterceptor, HttpErrorResponse, HttpResponse
} from '@angular/common/http';
import { Observable } from 'rxjs';
import {finalize, map, tap} from "rxjs/operators";
import {AjaxLoadNotifierService} from "../directive/ajax-load-notifier.service";

@Injectable({
  providedIn: 'root'
})
export class HttpAjaxBusyInterceptor implements HttpInterceptor {

  requestCounter = 0;

  constructor(private alnd: AjaxLoadNotifierService) {}

  intercept(request: HttpRequest<unknown>, next: HttpHandler): Observable<HttpEvent<unknown>> {
    this.beginRequest();
    return next.handle(request).pipe(
      finalize(() => {
        this.endRequest();
      })
    );
  }

  beginRequest() {
    this.requestCounter = Math.max(this.requestCounter, 0) + 1;
    if (this.requestCounter === 1) {
      this.alnd.busy.next(true);
    }
  }

  endRequest() {
    this.requestCounter = Math.max(this.requestCounter, 1) - 1;
    if (this.requestCounter === 0) {
      this.alnd.busy.next(false);
    }
  }
}
