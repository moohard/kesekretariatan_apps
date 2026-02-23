"use client"

import * as React from "react"
import { cn } from "../../lib/utils"
import { Button } from "./button"
import { ChevronLeft, ChevronRight, Check } from "lucide-react"

// Types
export interface Step {
  id: string
  title: string
  description?: string
  icon?: React.ReactNode
  isCompleted?: boolean
  isOptional?: boolean
}

export interface StepWizardProps {
  steps: Step[]
  currentStep: number
  onStepChange?: (step: number) => void
  allowStepClick?: boolean
  showStepNumbers?: boolean
  orientation?: "horizontal" | "vertical"
  className?: string
  children?: React.ReactNode
}

// Step Indicator Component
interface StepIndicatorProps {
  step: Step
  index: number
  currentStep: number
  totalSteps: number
  orientation: "horizontal" | "vertical"
  showStepNumbers: boolean
  allowStepClick: boolean
  onStepClick: (index: number) => void
}

function StepIndicator({
  step,
  index,
  currentStep,
  totalSteps,
  orientation,
  showStepNumbers,
  allowStepClick,
  onStepClick,
}: StepIndicatorProps) {
  const isCompleted = step.isCompleted || index < currentStep
  const isActive = index === currentStep
  const isLast = index === totalSteps - 1

  const handleClick = () => {
    if (allowStepClick && (isCompleted || index === currentStep + 1)) {
      onStepClick(index)
    }
  }

  return (
    <div
      className={cn(
        "flex",
        orientation === "horizontal" ? "flex-col items-center" : "items-start",
        { "cursor-pointer": allowStepClick && (isCompleted || index === currentStep + 1) }
      )}
      onClick={handleClick}
    >
      <div className="flex items-center">
        {/* Step Circle */}
        <div
          className={cn(
            "flex h-10 w-10 items-center justify-center rounded-full border-2 transition-all duration-200",
            isCompleted
              ? "border-primary bg-primary text-primary-foreground"
              : isActive
              ? "border-primary bg-background text-primary"
              : "border-muted bg-background text-muted-foreground",
            allowStepClick && (isCompleted || index === currentStep + 1) && "hover:border-primary/80"
          )}
        >
          {isCompleted ? (
            <Check className="h-5 w-5" />
          ) : step.icon ? (
            step.icon
          ) : showStepNumbers ? (
            <span className="text-sm font-semibold">{index + 1}</span>
          ) : (
            <span className="text-sm font-semibold">{index + 1}</span>
          )}
        </div>

        {/* Connector Line */}
        {orientation === "horizontal" && !isLast && (
          <div
            className={cn(
              "mx-2 h-0.5 w-12 transition-all duration-200",
              isCompleted ? "bg-primary" : "bg-muted"
            )}
          />
        )}
      </div>

      {/* Step Content */}
      <div
        className={cn(
          orientation === "horizontal" ? "mt-2 text-center" : "ml-4"
        )}
      >
        <p
          className={cn(
            "text-sm font-medium",
            isActive ? "text-foreground" : "text-muted-foreground"
          )}
        >
          {step.title}
          {step.isOptional && (
            <span className="ml-1 text-xs text-muted-foreground">(Opsional)</span>
          )}
        </p>
        {step.description && orientation === "vertical" && (
          <p className="mt-1 text-xs text-muted-foreground">{step.description}</p>
        )}
      </div>

      {/* Vertical Connector Line */}
      {orientation === "vertical" && !isLast && (
        <div
          className={cn(
            "ml-5 h-8 w-0.5 transition-all duration-200",
            isCompleted ? "bg-primary" : "bg-muted"
          )}
        />
      )}
    </div>
  )
}

// Main StepWizard Component
export function StepWizard({
  steps,
  currentStep,
  onStepChange,
  allowStepClick = false,
  showStepNumbers = true,
  orientation = "horizontal",
  className,
  children,
}: StepWizardProps) {
  const handleStepClick = (index: number) => {
    if (onStepChange) {
      onStepChange(index)
    }
  }

  return (
    <div className={cn("w-full", className)}>
      {/* Step Indicators */}
      <div
        className={cn(
          "flex",
          orientation === "horizontal"
            ? "flex-row items-start justify-between"
            : "flex-col"
        )}
      >
        {steps.map((step, index) => (
          <StepIndicator
            key={step.id}
            step={step}
            index={index}
            currentStep={currentStep}
            totalSteps={steps.length}
            orientation={orientation}
            showStepNumbers={showStepNumbers}
            allowStepClick={allowStepClick}
            onStepClick={handleStepClick}
          />
        ))}
      </div>

      {/* Step Content */}
      {children && (
        <div
          className={cn(
            "mt-8",
            orientation === "horizontal" ? "border-t pt-6" : "border-l pl-6"
          )}
        >
          {children}
        </div>
      )}
    </div>
  )
}

// Step Content Wrapper
export interface StepContentProps {
  children: React.ReactNode
  className?: string
}

export function StepContent({ children, className }: StepContentProps) {
  return <div className={cn("", className)}>{children}</div>
}

// Step Actions (Navigation Buttons)
export interface StepActionsProps {
  onPrevious?: () => void
  onNext?: () => void
  onSubmit?: () => void
  isPreviousDisabled?: boolean
  isNextDisabled?: boolean
  isSubmitting?: boolean
  previousLabel?: string
  nextLabel?: string
  submitLabel?: string
  currentStep: number
  totalSteps: number
  className?: string
}

export function StepActions({
  onPrevious,
  onNext,
  onSubmit,
  isPreviousDisabled = false,
  isNextDisabled = false,
  isSubmitting = false,
  previousLabel = "Sebelumnya",
  nextLabel = "Selanjutnya",
  submitLabel = "Simpan",
  currentStep,
  totalSteps,
  className,
}: StepActionsProps) {
  const isFirstStep = currentStep === 0
  const isLastStep = currentStep === totalSteps - 1

  return (
    <div className={cn("flex justify-between pt-4", className)}>
      <Button
        type="button"
        variant="outline"
        onClick={onPrevious}
        disabled={isFirstStep || isPreviousDisabled || isSubmitting}
        className={cn({ "invisible": isFirstStep })}
      >
        <ChevronLeft className="mr-2 h-4 w-4" />
        {previousLabel}
      </Button>

      {isLastStep ? (
        <Button
          type="button"
          onClick={onSubmit}
          disabled={isNextDisabled || isSubmitting}
        >
          {isSubmitting ? "Menyimpan..." : submitLabel}
        </Button>
      ) : (
        <Button
          type="button"
          onClick={onNext}
          disabled={isNextDisabled || isSubmitting}
        >
          {nextLabel}
          <ChevronRight className="ml-2 h-4 w-4" />
        </Button>
      )}
    </div>
  )
}

// Hook for managing step state
export function useStepWizard(totalSteps: number, initialStep = 0) {
  const [currentStep, setCurrentStep] = React.useState(initialStep)
  const [completedSteps, setCompletedSteps] = React.useState<Set<number>>(
    new Set()
  )

  const nextStep = React.useCallback(() => {
    if (currentStep < totalSteps - 1) {
      setCompletedSteps((prev) => new Set(prev).add(currentStep))
      setCurrentStep((prev) => prev + 1)
    }
  }, [currentStep, totalSteps])

  const previousStep = React.useCallback(() => {
    if (currentStep > 0) {
      setCurrentStep((prev) => prev - 1)
    }
  }, [currentStep])

  const goToStep = React.useCallback(
    (step: number) => {
      if (step >= 0 && step < totalSteps) {
        setCurrentStep(step)
      }
    },
    [totalSteps]
  )

  const reset = React.useCallback(() => {
    setCurrentStep(initialStep)
    setCompletedSteps(new Set())
  }, [initialStep])

  const isFirstStep = currentStep === 0
  const isLastStep = currentStep === totalSteps - 1

  return {
    currentStep,
    completedSteps,
    nextStep,
    previousStep,
    goToStep,
    reset,
    isFirstStep,
    isLastStep,
    totalSteps,
  }
}

export default StepWizard
