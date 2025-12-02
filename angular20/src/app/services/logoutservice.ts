import { Injectable, inject } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import {  Observable } from 'rxjs';

@Injectable({
  providedIn: 'root',
})

export class Logoutservice {
  private http = inject(HttpClient);
  private authUrl = "https://127.0.0.1:5000/app_logout"

  public logoutUser(): Observable<any> {
      return this.http.post(this.authUrl, null);
  }

}
