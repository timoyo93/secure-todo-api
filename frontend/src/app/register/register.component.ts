import { Component, OnInit } from '@angular/core';
import { AuthService } from '../services/auth.service';
import { Router } from '@angular/router';
import { HttpErrorResponse } from '@angular/common/http';

@Component({
  selector: 'app-register',
  templateUrl: './register.component.html',
  styleUrls: ['./register.component.scss'],
})
export class RegisterComponent implements OnInit {
  formValid: boolean = true;
  loading: boolean = false;
  userAlreadyTaken: boolean = false;

  constructor(private authService: AuthService, private router: Router) {}

  ngOnInit(): void {}

  registerUserSubmit(data: {
    username: string;
    password: string;
    checkPassword: string;
  }): void {
    this.loading = true;
    if (data.password !== data.checkPassword) {
      this.formValid = false;
      this.loading = false;
      return;
    }
    const request = { username: data.username, password: data.password };
    this.authService.registerUser(request).subscribe({
      next: () => {
        this.authService.loginUser(request).subscribe({
          next: () => {
            this.loading = false;
            this.authService.isAuthenticated.next(true);
            this.router.navigate(['/']).then();
          },
        });
      },
      error: (err) => {
        const error = err as HttpErrorResponse;
        if (error.status === 400) {
          this.userAlreadyTaken = true;
          this.loading = false;
          return;
        }
        this.loading = false;
        return;
      },
    });
  }
}
