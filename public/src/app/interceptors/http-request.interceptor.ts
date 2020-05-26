import { Injectable } from '@angular/core';
import {
  HttpRequest,
  HttpHandler,
  HttpEvent,
  HttpInterceptor, HttpErrorResponse, HttpResponse
} from '@angular/common/http';
import { Observable } from 'rxjs';
import {map, tap} from "rxjs/operators";

@Injectable()
export class HttpRequestInterceptor implements HttpInterceptor {

  constructor() {}

  intercept(request: HttpRequest<unknown>, next: HttpHandler): Observable<HttpEvent<unknown>> {
    console.log(request);
    const startTime = (new Date()).getTime();
    return next.handle(request).pipe(
      map(event => {
        if (event instanceof HttpResponse) {
          let endTime = (new Date()).getTime();
          let diff = endTime - startTime;
          console.log(`request successfully for host ${event.url}, elapsed time of ${diff}ms`)
        }
        return event;
      }),
      tap(event => {

      }, error => {
        if (error instanceof HttpErrorResponse) {
          let endTime = (new Date()).getTime();
          let diff = endTime - startTime;
          console.log(`failure to make request to host ${error.url}, elapsed time of ${diff}ms`)
        }
      })
    );
  }
}
